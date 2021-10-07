CREATE TABLE category
(
    id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    data json[],
    CONSTRAINT category_pkey PRIMARY KEY (id)

);

CREATE TABLE channel
(
    id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    count integer,
    country character varying(25) COLLATE pg_catalog."default",
    customurl character varying(250) COLLATE pg_catalog."default",
    description character varying COLLATE pg_catalog."default",
    favorites character varying COLLATE pg_catalog."default",
    highthumbnail character varying(255) COLLATE pg_catalog."default",
    itemcount integer,
    likes character varying COLLATE pg_catalog."default",
    localizeddescription character varying COLLATE pg_catalog."default",
    localizedtitle character varying(255) COLLATE pg_catalog."default",
    mediumthumbnail character varying(255) COLLATE pg_catalog."default",
    playlistcount integer,
    playlistitemcount integer,
    playlistvideocount integer,
    playlistvideoitemcount integer,
    publishedat timestamp with time zone,
    thumbnail character varying(255) COLLATE pg_catalog."default",
    lastupload timestamp with time zone,
    title character varying(255) COLLATE pg_catalog."default",
    uploads character varying(100) COLLATE pg_catalog."default",
    channels character varying[] COLLATE pg_catalog."default",
    CONSTRAINT channel_pkey PRIMARY KEY (id)
);

CREATE TABLE channelsync
(
    id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    synctime timestamp with time zone,
    uploads character varying(100) COLLATE pg_catalog."default",
    CONSTRAINT channelsync_pkey PRIMARY KEY (id)

);

CREATE TABLE playlist
(
    id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    channelid character varying(100) COLLATE pg_catalog."default",
    channeltitle character varying(255) COLLATE pg_catalog."default",
    count integer,
    itemcount integer,
    description character varying COLLATE pg_catalog."default",
    highthumbnail character varying(255) COLLATE pg_catalog."default",
    localizeddescription character varying COLLATE pg_catalog."default",
    localizedtitle character varying(255) COLLATE pg_catalog."default",
    maxresthumbnail character varying(255) COLLATE pg_catalog."default",
    mediumthumbnail character varying(255) COLLATE pg_catalog."default",
    publishedat timestamp with time zone,
    standardthumbnail character varying(255) COLLATE pg_catalog."default",
    thumbnail character varying(255) COLLATE pg_catalog."default",
    title character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT playlist_pkey PRIMARY KEY (id)
);

CREATE TABLE playlistvideo
(
    id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    videos character varying[] COLLATE pg_catalog."default",
    CONSTRAINT playlistvideo_pkey PRIMARY KEY (id)

);

CREATE TABLE video
(
    id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    caption character varying(255) COLLATE pg_catalog."default",
    categoryid character varying(20) COLLATE pg_catalog."default",
    channelid character varying(100) COLLATE pg_catalog."default",
    channeltitle character varying(255) COLLATE pg_catalog."default",
    defaultaudiolanguage character varying(255) COLLATE pg_catalog."default",
    defaultlanguage character varying(255) COLLATE pg_catalog."default",
    definition smallint,
    description character varying COLLATE pg_catalog."default",
    dimension character varying(20) COLLATE pg_catalog."default",
    duration bigint,
    highthumbnail character varying(255) COLLATE pg_catalog."default",
    licensedcontent boolean,
    livebroadcastcontent character varying(255) COLLATE pg_catalog."default",
    localizeddescription character varying COLLATE pg_catalog."default",
    localizedtitle character varying(255) COLLATE pg_catalog."default",
    maxresthumbnail character varying(255) COLLATE pg_catalog."default",
    mediumthumbnail character varying(255) COLLATE pg_catalog."default",
    projection character varying(255) COLLATE pg_catalog."default",
    publishedat timestamp with time zone,
    standardthumbnail character varying(255) COLLATE pg_catalog."default",
    tags character varying[] COLLATE pg_catalog."default",
    thumbnail character varying(255) COLLATE pg_catalog."default",
    title character varying(255) COLLATE pg_catalog."default",
    blockedregions character varying(100)[] COLLATE pg_catalog."default",
    allowedregions character varying(100)[] COLLATE pg_catalog."default",
    CONSTRAINT video_pkey PRIMARY KEY (id)
);
