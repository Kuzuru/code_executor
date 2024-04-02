import asyncio
import logging
import os
import resource
import subprocess
import aiohttp
from types import SimpleNamespace

from aiohttp import web, ClientSession

from tempfile_helper import TempFileManager

AGENT_PORT = int(os.getenv("AGENT_PORT", "3030"))
TIMEOUT = int(os.getenv("TIMEOUT", "10"))
TESTS_PATH = "/home/student/tests/"


async def health_check_handler(request: web.Request) -> web.Response:
    return web.json_response({})


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
    duration = 0

    async with run_lock:
        with TempFileManager(directory=TESTS_PATH, files=files) as manager:
            try:
                # Recording process duration
                usage_start = resource.getrusage(resource.RUSAGE_CHILDREN)

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

                usage_end = resource.getrusage(resource.RUSAGE_CHILDREN)

                # Cast seconds to milliseconds
                duration = (usage_end.ru_utime - usage_start.ru_utime) * 1000

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
        "duration": round(duration, 2),
    }

    return web.json_response(result)


async def get_available_programming_languages(request: web.Request) -> web.Response:
    result = {
        "languages": [
            "Go 1.19",
            "Python 3.11",
            "Java"
        ]
    }

    return web.json_response(result)


async def register_new_user(request: web.Request) -> web.Response:
    body = await request.json()

    if "name" not in body or "password" not in body:
        result = {
            "state": 1,
            "success": False,
            "error": True,
            "message": "Missing required fields: name, password"
        }

        return web.json_response(result, status=400)

    async with aiohttp.ClientSession() as session:
        async with session.post('http://localhost:8080/user/register', json=body) as response:
            response_json = await response.json()
            if response.status != 200:
                result = {
                    "state": response_json["state"],
                    "success": response_json["success"],
                    "error": response_json["error"],
                    "message": response_json["message"]
                }

                return web.json_response(result, status=response.status)

    result = {
        "state": 0,
        "success": True,
        "error": False,
        "message": "ok"
    }

    return web.json_response(result)


def setup_routes(http: web.Application) -> None:
    http.router.add_post("/run", run)
    http.router.add_post("/user/register", register_new_user)
    http.router.add_get("/get_available_langs", get_available_programming_languages)
    http.router.add_get("/healthcheck", health_check_handler)


app = web.Application()
run_lock = asyncio.Lock()
setup_routes(app)
logging.basicConfig(level=logging.DEBUG)

web.run_app(
    app,
    host="0.0.0.0",
    port=AGENT_PORT,
)
