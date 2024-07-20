# BrewScylla
This is a coffee house backend for [the frontend app]().

## Design Choices
- Using **ScyllaDB** as a proof of concept (PoC) to explore its capabilities and performance. 
- **Elasticsearch** is used for search capabilities, allowing users to perform fast and flexible searches on product names, including partial and fuzzy matches.
- The web application is buuilt with **Go and the Gin framework** to handle requests.

## DB schema
Although ScyllaDB is a NoSQL database, it operates as a wide-column store. In this model, the schema requires that columns be pre-defined. This means that, unlike some other NoSQL databases that offer flexible or dynamic schemas, the structure of each table in ScyllaDB is fixed and must be defined at creation time.

The schema used in this project is:

![db schema image](images/DB%20schema.png)

Some important considerations while designing wide-column NoSQL DBs
1. Filtering can only be done on Primary columns or Clustering columns.
2. Primary keys are always cumpolsary while querying, but clustering keys can be optional.

    1. If any clustering columns are used in a query, you must specify all previous clustering columns (from left to right) with equality comparisons (=) before using any inequality comparisons (<, >, <=, >=, IN).
    2. Clustering columns can be skipped from right to left. This means you can use fewer clustering columns than defined in the primary key, but you must start from the leftmost clustering column and cannot skip columns in between.

Credits to DataStax Data Modelling [documentation](https://www.datastax.com/dev/modeling) for helping with data modelling best practices.

# TODOs
1. Making sure that phone numbers are unique in user details table
2. Inserting order and updating status. Making sure that product exists/user exists before making order. (DONE)
3. Handling coins. (DONE)
4. Add discount to coupon code table/entity. (DONE)
4. Taking care of packs and continuous buying items in packs
5. Search products
6. Handling magic-link login
7. Use UUIDs not uuid as string
7. Admin APIs