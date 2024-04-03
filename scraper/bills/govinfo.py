from datetime import datetime
import logging
import urllib.parse
from typing import Any, TypeVar
from bs4 import Tag
from .fetch import Fetch
from .common import BASE_DATE, key_or_default
from .bill import Bill, Action

T = TypeVar("T")


class GovInfoBillsResponsePackage:
    def __init__(self) -> None:
        self.id = ""
        self.date_issued = BASE_DATE
        self.last_modified = BASE_DATE
        self.link = ""
        self.title = ""


class GovInfoBillsResponse:
    def __init__(self) -> None:
        self.count = 0
        self.next_page = ""
        self.packages: list[GovInfoBillsResponsePackage] = []


class GovInfoBills:
    def __init__(self, api_key: str, fetch: Fetch) -> None:
        self.api_key = api_key
        self.fetch = fetch
        self.url_base = "https://api.govinfo.gov"

    def __format_datetime(self, d: datetime) -> str:
        date = d.strftime("%Y-%m-%d")
        time = d.strftime("%H:%M:%S")
        return f"{date}T{time}Z"

    def __url(self, path: str, query: dict[str, Any]) -> str:
        return self.url_base + path + "?" + urllib.parse.urlencode(query)

    def __transform_package(self, p: Any) -> GovInfoBillsResponsePackage:
        r = GovInfoBillsResponsePackage()
        r.id = key_or_default(p, "packageId", "")
        r.link = key_or_default(p, "packageLink", "")
        r.title = key_or_default(p, "title", "")

        issued = key_or_default(p, "dateIssued", "")
        if issued != "":
            r.date_issued = datetime.strptime(issued, "%Y-%m-%d")

        modified = key_or_default(p, "lastModified", "")
        if modified != "":
            r.last_modified = datetime.strptime(modified, "%Y-%m-%dT%H:%M:%SZ")

        return r

    def __transform_response(self, response: Any) -> GovInfoBillsResponse:
        r = GovInfoBillsResponse()
        r.count = key_or_default(response, "count", 0)
        r.next_page = key_or_default(response, "nextPage", "")

        packages = key_or_default(response, "packages", [])
        r.packages = [self.__transform_package(p) for p in packages]

        return r

    def fetch_bills(
        self, from_date: datetime, limit: int = 1000, offset: int = 0
    ) -> GovInfoBillsResponse:
        url_path = f"/collections/BILLS/{self.__format_datetime(from_date)}"
        url_query = {"pageSize": limit, "offset": offset}
        url = self.__url(url_path, url_query)
        response = self.fetch.json_request(url, headers={"X-Api-Key": self.api_key})
        return self.__transform_response(response)


class GovInfoBill:
    def __init__(self, api_key: str, fetch: Fetch) -> None:
        self.api_key = api_key
        self.fetch = fetch

    def __transform_action(self, t: Tag) -> Action | None:
        a = Action()

        date = t.find("actionDate")
        if isinstance(date, Tag):
            a.date = datetime.strptime(date.text, "%Y-%m-%d")

        text = t.find("text")
        if isinstance(text, Tag):
            a.text = text.text
        else:
            return None

        return a

    def __fetch_actions(self, url: str) -> list[Action]:
        response = self.fetch.xml_request(url)

        action_item_tags: list[Tag] = []

        actions_tag = response.find("actions")
        if isinstance(actions_tag, Tag):
            action_item_tags = actions_tag.find_all("item")

        data: dict[str, Action] = {}
        for action_item_tag in action_item_tags:
            action = self.__transform_action(action_item_tag)
            if action is not None:
                key = f"{action.text}::{action.date.strftime('%Y-%m-%dT%H:%M:%S')}"
                data[key] = action

        actions = list(data.values())
        actions.sort(key=lambda a: a.date, reverse=True)
        return actions

    def __fetch_text(self, url: str) -> str:
        return self.fetch.html_request(url, headers={"X-Api-Key": self.api_key}).text

    def __transform_response(self, response: Any) -> Bill:
        b = Bill()
        b.bill_id = key_or_default(response, "packageId", "")
        b.title = key_or_default(response, "title", "")
        b.url = key_or_default(response, "detailsLink", "")
        b.version = key_or_default(response, "billVersion", "")
        b.type = key_or_default(response, "billType", "")

        issued = key_or_default(response, "dateIssued", "")
        if issued != "":
            b.issued = datetime.strptime(issued, "%Y-%m-%d")

        updated = key_or_default(response, "lastModified", "")
        if updated != "":
            b.updated = datetime.strptime(updated, "%Y-%m-%dT%H:%M:%SZ")

        bill_status_link = ""
        related_links = key_or_default(response, "related", "")
        if related_links != "":
            bill_status_link = key_or_default(related_links, "billStatusLink", "")
        if bill_status_link != "":
            logging.info("Fetching actions for bill:", b.bill_id)
            b.actions = self.__fetch_actions(bill_status_link)

        bill_text_link = ""
        downloads = key_or_default(response, "download", [])
        if downloads != []:
            bill_text_link = key_or_default(downloads, "txtLink", "")
        if bill_text_link != "":
            logging.info("Fetching text for bill:", b.bill_id)
            b.text = self.__fetch_text(bill_text_link)

        return b

    def fetch_bill(self, url: str) -> Bill:
        logging.info(f"Fetching info for bill: {url}")
        response = self.fetch.json_request(url, headers={"X-Api-Key": self.api_key})
        return self.__transform_response(response)
