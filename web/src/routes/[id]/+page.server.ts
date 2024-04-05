import { billValidator, bills } from "$lib";
import { error } from "@sveltejs/kit";
import type { RequestEvent } from "../$types";

export async function load({ params }: RequestEvent) {
    let id = ""
    if ("id" in params && typeof params.id === "string") id = params.id

    const doc = await bills.getBill(id)
    if (doc === null) error(404, "Not found")

    const bill = billValidator.parse(doc)
    return {
        bill
    }
}