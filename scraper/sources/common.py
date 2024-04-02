from datetime import datetime
from typing import Any, TypeVar

BASE_DATE = datetime(1900, 0, 0)

T = TypeVar("T")


def key_or_default(d: Any, key: str, default: T) -> T:
    if key in d:
        return d[key]
    return default
