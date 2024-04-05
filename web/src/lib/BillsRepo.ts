import { Collection, MongoClient, type Document } from "mongodb";

export default class BillsRepo {
    protected client: MongoClient
    protected dbName: string
    protected colName: string
    protected hasConnected: boolean

    constructor(mongoUrl: string, dbName: string, colName: string) {
        this.client = new MongoClient(mongoUrl)
        this.dbName = dbName
        this.colName = colName
        this.hasConnected = false
    }

    protected getClient = async (): Promise<MongoClient> => {
        if (!this.hasConnected) await this.client.connect()
        return this.client
    }

    protected getCollection = async (): Promise<Collection<Document>> => {
        const client = await this.getClient()
        const db = client.db(this.dbName)
        const col = db.collection(this.colName)
        return col
    }

    getBills = async (args?: {
        limit?: number,
        offset?: number
    }) => {
        const col = await this.getCollection()
        const docs = await col.find().skip(args?.offset ?? 0).limit(args?.limit ?? 20).sort({ "updated": -1 }).project({ text: 0 }).toArray()
        return docs
    }

    getBill = async (id: string) => {
        const col = await this.getCollection()
        const doc = await col.findOne({ bill_id: id })
        return doc
    }
}