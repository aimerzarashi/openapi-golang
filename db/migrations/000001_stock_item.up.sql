-- CreateTable
CREATE TABLE stock_item (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL,

    CONSTRAINT stock_item_pkey PRIMARY KEY ("id")
);
