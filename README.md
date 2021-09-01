# international-pricebook-challenge

- Basically the ideia of the challenge is to create an API that is capable to handle products with different prices based on different currencies.

## Database Structure

The postgres database was used to store all the informations about the products and its prices as the following:

### Tables:
- Product
- Price
- Currency
- Country

### Relations:
- A product may have several prices
- A price can only be associated to 1 product and 1 currency
- A currency can be associated to many prices and many countries

### Entity Relationship Diagram

![Database Relations](./images/DBDiagram.png)