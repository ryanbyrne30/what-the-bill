import logging
import requests
import time
import random
from typing import Any
from bs4 import BeautifulSoup
from .user_agents import random_user_agent


class Fetch:
    def __init__(
        self,
        timeout_min: int = 2,
        timeout_max: int = 10,
    ) -> None:
        self.last_request_time: float = 0
        self.timeout_min = timeout_min
        self.timeout_max = timeout_max
        self.user_agent = random_user_agent()
        self.request_count = 0

    def __cycle_agent(self):
        if self.request_count > 10:
            self.user_agent = random_user_agent()
            self.request_count = 0

    def __create_session(
        self, headers: dict[str, str | bytes] = {}
    ) -> requests.Session:
        h: dict[str, str | bytes] = {"User-Agent": self.user_agent}
        h.update(headers)

        sess = requests.Session()
        sess.headers = h
        return sess

    def sleep(self):
        delta = time.time() - self.last_request_time
        if delta < self.timeout_min:
            sleep_time_ms = random.randrange(0, 1000) / 1000
            sleep_time_sec = random.randrange(self.timeout_min, self.timeout_max)
            sleep_time = sleep_time_sec + sleep_time_ms
            logging.info(f"Sleeping for {round(sleep_time, 2)} seconds...")
            time.sleep(sleep_time)
        self.last_request_time = time.time()

    def __simple_request(
        self, url: str, headers: dict[str, str | bytes], count: int = 0
    ) -> requests.Response | None:
        logging.debug(f"Sending simple request to {url}")
        max_count = 10
        self.sleep()
        self.__cycle_agent()
        response = self.__create_session(headers).get(url)
        if response.status_code == 200:
            return response
        logging.warn(f"Received status code {response.status_code} for url {url}")
        if count >= max_count:
            return None
        else:
            time.sleep(5)
        return self.__simple_request(url, headers, count + 1)

    def text_request(self, url: str, headers: dict[str, str | bytes] = {}) -> str:
        response = self.__simple_request(url, headers)
        if response is not None:
            return response.text
        return ""

    def html_request(
        self, url: str, headers: dict[str, str | bytes] = {}
    ) -> BeautifulSoup:
        response = self.__simple_request(url, headers)
        if response is not None:
            return BeautifulSoup(response.content, "html.parser")
        return BeautifulSoup("", "html.parser")

    def xml_request(
        self, url: str, headers: dict[str, str | bytes] = {}
    ) -> BeautifulSoup:
        response = self.__simple_request(url, headers)
        if response is not None:
            return BeautifulSoup(response.content, "xml")
        return BeautifulSoup("", "xml")

    def json_request(self, url: str, headers: dict[str, str | bytes] = {}) -> Any:
        response = self.__simple_request(url, headers)
        if response is not None:
            return response.json()
        return {}
