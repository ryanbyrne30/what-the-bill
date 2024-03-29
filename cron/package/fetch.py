from seleniumwire import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.chrome.options import Options
import requests 
from .user_agents import random_user_agent
import time
import random

LAST_REQUEST_TIME = 0

class Fetch:
    def __init__(
            self, 
            timeout_min: int = 2, 
            timeout_max: int = 10) -> None:
        self.last_request_time = 0
        self.session = self.__create_session() 
        self.timeout_min = timeout_min
        self.timeout_max = timeout_max
        self.webdriver = self.__setup_webdriver()

    def __create_session(self) -> requests.Session:
        headers = {
            'User-Agent': random_user_agent()
        }
        sess =  requests.Session()
        sess.headers = headers
        return sess
    
    def sleep(self):
        delta = time.time() - self.last_request_time 
        if delta < self.timeout_min:
            sleep_time_ms = random.randrange(0, 1000) / 1000
            sleep_time_sec = random.randrange(self.timeout_min, self.timeout_max)
            sleep_time = sleep_time_sec + sleep_time_ms
            print(f"Sleeping for {sleep_time} seconds...")
            time.sleep(sleep_time)
        self.last_request_time = time.time()

    def simple_request( self, url: str) -> str:
        self.sleep()
        print("Sending simple request to:", url)
        response = self.session.get(url)
        return response.content

    def __intercept_request(self, req):
        req.headers['User-Agent'] = random_user_agent()

    def __setup_webdriver(self) -> webdriver.Chrome:
        options = Options()
        options.add_argument('--headless')
        driver = webdriver.Chrome(options=options)
        driver.request_interceptor = self.__intercept_request
        return driver

    
    def dynamic_request(self, url: str, wait_for: str | None):
        self.sleep()
        print("Sending dynamic request to:", url)
        self.webdriver.get(url)
        html = ""

        try:
            if wait_for:
                WebDriverWait(self.webdriver, 10).until(
                    EC.presence_of_element_located((By.CSS_SELECTOR, wait_for))
                )
            html = self.webdriver.page_source
        finally:
            self.webdriver.quit()

        return html

