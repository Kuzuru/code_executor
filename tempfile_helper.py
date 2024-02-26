import os
import shutil
import tempfile
from typing import List, Optional, Union, Type, Tuple
from types import SimpleNamespace, TracebackType


class TempFileManager:
    def __init__(self, directory: str, files: List[SimpleNamespace], cleanup: bool = True):
        self.directory = directory
        self.files = files
        self.cleanup = cleanup

    def __enter__(self) -> 'TempFileManager':
        self.temp_dir = tempfile.mkdtemp()
        for file in self.files:
            file_path = os.path.join(self.temp_dir, file.name)
            with open(file_path, 'wb') as f:
                f.write(file.content)
        return self

    def __exit__(self, exc_type: Optional[Union[Type[BaseException], Tuple[Type[BaseException], ...]]],
                 exc_value: Optional[BaseException],
                 traceback: Optional[TracebackType]) -> None:
        if self.cleanup:
            shutil.rmtree(self.temp_dir)

    def cleanup_files(self) -> None:
        shutil.rmtree(self.temp_dir)

    def get_temp_dir_path(self) -> str:
        return self.temp_dir
