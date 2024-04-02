import scraper.sources.common as common


class Action:
    def __init__(self) -> None:
        self.date = common.BASE_DATE
        self.text = ""


class Bill:
    def __init__(self) -> None:
        self.id = ""
        self.bill_id = ""
        self.title = ""
        self.url = ""
        self.text = ""
        self.version = ""
        self.type = ""
        self.issued = common.BASE_DATE
        self.updated = common.BASE_DATE
        self.actions: list[Action] = []
