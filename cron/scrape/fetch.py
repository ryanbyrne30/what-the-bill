import logging
from seleniumwire import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service as ChromeService
from webdriver_manager.chrome import ChromeDriverManager
import requests
from .user_agents import random_user_agent
from .proxy import Proxy
import time
import random


class Fetch:
    def __init__(
        self,
        proxy: Proxy | None = None,
        timeout_min: int = 2,
        timeout_max: int = 10,
    ) -> None:
        self.last_request_time: float = 0
        self.timeout_min = timeout_min
        self.timeout_max = timeout_max
        self.user_agent = random_user_agent()
        self.request_count = 0
        self.proxy = proxy

    def __log_ip(self):
        content = self.simple_request("http://checkip.amazonaws.com/")
        logging.info(f"IP: {str(content)}")

    def __cycle_agent(self):
        if self.request_count > 10:
            self.user_agent = random_user_agent()
            self.request_count = 0

    def __create_session(self) -> requests.Session:
        headers: dict[str, str | bytes] = {"User-Agent": self.user_agent}
        sess = requests.Session()
        sess.headers = headers
        if self.proxy is not None:
            sess.proxies = {
                "http": self.proxy.socks_connection(),
                "https": self.proxy.socks_connection(),
            }
        return sess

    def sleep(self):
        delta = time.time() - self.last_request_time
        if delta < self.timeout_min:
            sleep_time_ms = random.randrange(0, 1000) / 1000
            sleep_time_sec = random.randrange(self.timeout_min, self.timeout_max)
            sleep_time = sleep_time_sec + sleep_time_ms
            # print(f"Sleeping for {sleep_time} seconds...")
            time.sleep(sleep_time)
        self.last_request_time = time.time()

    def simple_request(self, url: str, count: int = 0) -> bytes:
        logging.debug(f"Sending simple request to {url}")
        max_count = 10
        self.sleep()
        self.__cycle_agent()
        response = self.__create_session().get(url)
        if response.status_code == 200:
            return response.content
        logging.warn(f"Received status code: {response.status_code}")
        if count >= max_count:
            return bytes("", "utf-8")
        if self.proxy is not None:
            logging.error("Resetting proxy")
            self.proxy.reset()
            time.sleep(10)
            self.__log_ip()
        else:
            time.sleep(5)
        return self.simple_request(url, count + 1)

    def __intercept_request(self, req):
        req.headers["User-Agent"] = self.user_agent

    def __setup_webdriver(self) -> webdriver.Chrome:
        options = Options()
        options.add_argument("--headless")
        options.add_argument("--log-level=4")
        options.add_experimental_option("excludeSwitches", ["enable-logging"])
        if self.proxy is not None:
            options.add_argument(f"--proxy-server={self.proxy.socks_connection()}")
        driver = webdriver.Chrome(
            service=ChromeService(ChromeDriverManager().install()), options=options
        )
        driver.request_interceptor = self.__intercept_request
        return driver

    def dynamic_request(self, url: str, wait_for: str | None, count: int = 0):
        max_count = 10
        self.sleep()
        self.__cycle_agent()

        webdriver = self.__setup_webdriver()
        webdriver.get(url)
        html = ""

        try:
            if wait_for:
                WebDriverWait(webdriver, 10).until(
                    EC.presence_of_element_located((By.CSS_SELECTOR, wait_for))
                )
            html = webdriver.page_source
        except:
            if count < max_count and self.proxy is not None:
                logging.error("Resetting proxy...")
                self.proxy.reset()
                time.sleep(10)
                self.__log_ip()
                return self.dynamic_request(url, wait_for, count=count + 1)
            return html
        finally:
            webdriver.quit()

        return html
