// place files you want to import through the `$lib` alias in this folder.

import { env } from "$env/dynamic/private";
import { z } from "zod";
import BillsRepo from "./BillsRepo";

export const bills = new BillsRepo(
    env.MONGODB_URL ?? "",
    env.MONGODB_DATABASE ?? "",
    env.MONGODB_COLLECTION ?? ""
)

export const billSummaryValidator = z.object({
    bill_id: z.string(),
    url: z.string(),
    title: z.string(),
    short_title: z.string(),
    updated: z.date(),
    issued: z.date(),
    actions: z.array(z.object({
        date: z.date(),
        text: z.string()
    }))
})
export type BillSummary = z.infer<typeof billSummaryValidator>
export const billsValidator = z.array(billSummaryValidator);

export const billValidator = z.object({
    bill_id: z.string(),
    title: z.string(),
    short_title: z.string(),
    url: z.string(),
    text: z.string(),
    updated: z.date(),
    issued: z.date(),
    actions: z.array(z.object({
        date: z.date(),
        text: z.string()
    }))
})
export type Bill = z.infer<typeof billValidator>