from pymongo import MongoClient, UpdateOne


class BillData:
    def __init__(self, bill_id: str, data: any) -> None:
        self.bill_id = bill_id
        self.data = data


class Mongo:
    def __init__(self, conn_str: str, db_name: str, col_name: str) -> None:
        self.client = MongoClient(conn_str)
        self.db = self.client[db_name]
        self.col = self.db[col_name]

    def upsert_bills(self, bill_data: list[BillData]):
        operations = [
            UpdateOne({"bill_id": d.bill_id}, {"$set": d.data}, upsert=True)
            for d in bill_data
        ]
        return self.col.bulk_write(operations)
