import asyncio
import logging
import os
import subprocess
from types import SimpleNamespace

from aiohttp import web

from tempfile_helper import TempFileManager

AGENT_PORT = int(os.getenv("PORT", "3000"))
TIMEOUT = int(os.getenv("TIMEOUT", "10"))
TESTS_PATH = "/home/student/tests/"


async def run(request: web.Request) -> web.Response:
    body = await request.json()

    files = [
        SimpleNamespace(name=f["name"], content=f["content"].encode()) for f in body["files"]
    ]

    timeout = False
    stdout = b""
    stderr = b""
    return_code = 1
    oom_killed = False

    async with run_lock:
        with TempFileManager(directory=TESTS_PATH, files=files) as manager:
            try:
                proc = subprocess.run(
                    (
                        f"chown -R student {manager.temp_dir} "
                        f"&& chown -R student /tmp/ "
                        f"&& chown -R student /home/student/ "
                        f"&& sudo -u student sh -c 'cd \"{manager.temp_dir}\" && {body['command']}'"
                    ),
                    capture_output=True,
                    timeout=TIMEOUT,
                    shell=True,
                )

                stdout = proc.stdout
                stderr = proc.stderr
                return_code = proc.returncode
            except subprocess.TimeoutExpired:
                timeout = True

    result = {
        "exit_code": return_code,
        "stdout": stdout.decode(),
        "stderr": stderr.decode(),
        "oom_killed": oom_killed,
        "timeout": timeout,
        "duration": 0,
    }

    return web.json_response(result)


async def get_available_programming_languages() -> web.Response:
    result = {
        "languages": [
            "Go 1.19",
            "Python 3.11",
            "Java"
        ]
    }

    return web.json_response(result)


def setup_routes(http: web.Application) -> None:
    http.router.add_post("/run", run)
    http.router.add_get("/get_available_langs", get_available_programming_languages)


app = web.Application()
run_lock = asyncio.Lock()
setup_routes(app)
logging.basicConfig(level=logging.DEBUG)

web.run_app(
    app,
    host="0.0.0.0",
    port=AGENT_PORT,
)
