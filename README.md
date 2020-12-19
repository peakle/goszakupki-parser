## Build or Install using make

* `make build` - builds the binary

The `Makefile` builds the binary and adds the short git commit id and the version to the final build.

## Usage

Service parse tenders from goszakupki website.

Purchase:
    To get parsed information you need query `/get/purchase` url.

For filter information you should use following json params:
##### `to_date`  - filter to parse date, equal to: <= 'your_date'. 
##### `from_date`  - filter from parse date, equal to: >= 'your_date'.
##### `customer` - filter by customer INN
##### `region` - filter by bidding region
##### `price_from` - filter by contract price bigger or equal to your_price
##### `price_to` - filter contract price less or equal to your_price
##### `grnt_share_from` - filter garantee share bigger or equal to your share
##### `grnt_share_to` - filter by garantee share less or equal to your share
