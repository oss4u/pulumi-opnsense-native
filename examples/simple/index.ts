import * as opnsense from "@pulumi/opnsense";

const page = new opnsense.StaticPage("page", {
    indexContent: "<html><body><p>Hello world!</p></body></html>",
});

export const bucket = page.bucket;
export const url = page.websiteUrl;
