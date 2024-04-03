from datetime import datetime
from typing import Any, TypeVar

BASE_DATE = datetime(1900, 1, 1)

T = TypeVar("T")


def key_or_default(d: Any, key: str, default: T) -> T:
    if key in d:
        return d[key]
    return default
