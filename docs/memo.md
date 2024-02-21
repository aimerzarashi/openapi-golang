```mermaid
flowchart TB
    sell/shop --> sell/item
    sell/category --> sell/item
    sell/category --> sell/shop
    sell/warehousing --> sell/item
    inbound/request --> sell/warehousing
    inbound/order --> inbound/request
    storage/stock --> inbound/order
    buy/display --> storage/stock
    buy/cart --> buy/display
    buy/payment --> buy/cart
    buy/delivery --> buy/payment
    outbound/request --> buy/delivery
    outbound/order --> outbound/request
    delivery/request --> buy/delivery
    delivery/order --> delivery/request
```