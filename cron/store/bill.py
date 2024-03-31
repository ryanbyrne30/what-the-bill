from typing import Any, TypeVar
from datetime import datetime
from . import config

T = TypeVar("T")


class Bill:
    def __init__(self) -> None:
        self.short_title = ""
        self.full_title = ""
        self.category = ""
        self.url = ""
        self.collection = ""
        self.sudoc_class_number = ""
        self.congress_number = ""
        self.congress_session = ""
        self.last_action_date: datetime = config.BASE_DATE
        self.action = ""
        self.actions = ""
        self.bill_number = ""
        self.bill_version = ""
        self.sponsors = ""
        self.cosponsors = ""
        self.committees = ""
        self.us_code_reference = ""
        self.text = ""
        self.updated_at: datetime = config.BASE_DATE
        self.congress_updated_at: datetime = config.BASE_DATE

        self.__dict__ = {
            "short_title": self.short_title,
            "full_title": self.full_title,
            "category": self.category,
            "url": self.url,
            "collection": self.collection,
            "sudoc_class_number": self.sudoc_class_number,
            "congress_number": self.congress_number,
            "congress_session": self.congress_session,
            "last_action_date": datetime.strftime(
                self.last_action_date, config.DATE_FMT
            ),
            "action": self.action,
            "actions": self.actions,
            "bill_number": self.bill_number,
            "bill_version": self.bill_version,
            "sponsors": self.sponsors,
            "cosponsors": self.cosponsors,
            "committees": self.committees,
            "us_code_reference": self.us_code_reference,
            "text": self.text,
            "updated_at": datetime.strftime(self.updated_at, config.DATE_FMT),
            "congress_updated_at": datetime.strftime(
                self.congress_updated_at, config.DATE_FMT
            ),
        }
