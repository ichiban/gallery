import type {LoadEvent} from "@sveltejs/kit";

export type pageData = App.PageData & {images?: Promise<Array<Image>>};

export interface Image {
    url: string,
}

export async function load({fetch}: LoadEvent): Promise<pageData> {
    const res = await fetch('/images');
    const images = res.json();

    return { images }
}