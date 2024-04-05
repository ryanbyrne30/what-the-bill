import { bills, billsValidator } from "$lib";
import type { RequestEvent } from "@sveltejs/kit";

export async function load({ url }: RequestEvent) {
    const page = Number(url.searchParams.get("page") ?? 0)
    const count = 10

    const docs = await bills.getBills({
        limit: count,
        offset: page * count
    })

    const result = billsValidator.parse(docs)

    return { bills: result, page }
}
