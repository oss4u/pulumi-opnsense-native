import * as opnsense from "@pulumi/opnsense";

const random = new opnsense.Random("my-random", { length: 24 });

export const output = random.result;