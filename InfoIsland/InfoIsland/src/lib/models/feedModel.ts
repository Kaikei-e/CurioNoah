export type Feed = {
    title: string;
    link: string;
    feedLink: string;
    links: string[];
    updated: string;
    updatedParsed: string;
    published: string;
    publishedParsed: string;
    language: string;
    image: Image;
    copyright: string;
    extensions: FeedExtensions;
    items: Item[];
    feedType: string;
    feedVersion: string;
}

export type FeedExtensions = {
    atom: PurpleAtom;
}

export type PurpleAtom = {
    link: PurpleLink[];
}

export type PurpleLink = {
    name: LinkName;
    value: string;
    attrs: PurpleAttrs;
    children: Children;
}

export type PurpleAttrs = {
    href: string;
    rel: string;
    type: string;
}

export type Children = {}

export enum LinkName {
    Link = "link",
}

export type Image = {
    url: string;
    title: string;
}

export type Item = {
    title: string;
    description?: string;
    link: string;
    links: string[];
    published: string;
    publishedParsed: string;
    author: Author;
    authors: Author[];
    guid: string;
    dcExt: DcEXT;
    extensions: ItemExtensions;
    categories?: string[];
}

export type Author = {
    name: string;
}

export type DcEXT = {
    creator: string[];
}

export type ItemExtensions = {
    atom: FluffyAtom;
    dc: Dc;
    media?: Media;
}

export type FluffyAtom = {
    link: FluffyLink[];
}

export type FluffyLink = {
    name: LinkName;
    value: string;
    attrs: FluffyAttrs;
    children: Children;
}

export type FluffyAttrs = {
    href: string;
    rel: Rel;
}

export enum Rel {
    Standout = "standout",
}

export type Dc = {
    creator: Creator[];
}

export type Creator = {
    name: CreatorName;
    value: string;
    attrs: Children;
    children: Children;
}

export enum CreatorName {
    Creator = "creator",
    Credit = "credit",
    Description = "description",
}

export type Media = {
    content: Content[];
    credit?: Creator[];
    description?: Creator[];
}

export type Content = {
    name: ContentName;
    value: string;
    attrs: ContentAttrs;
    children: Children;
}

export type ContentAttrs = {
    height: string;
    medium: Medium;
    url: string;
    width: string;
}

export enum Medium {
    Image = "image",
}

export enum ContentName {
    Content = "content",
}
