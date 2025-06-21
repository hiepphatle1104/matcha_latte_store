db = new Mongo().getDB("matcha_latte_db");

db.createCollection("products");
db.createCollection("orders");
db.createCollection("carts");
db.createCollection("transactions");
