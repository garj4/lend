package db

const createPeopleTableQuery = `CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)`
const createTransactionsTableQuery = `CREATE TABLE IF NOT EXISTS transactions (id INTEGER PRIMARY KEY, event TEXT, amount FLOAT, date DATETIME, person INTEGER)`

const addPersonQuery = `INSERT INTO people (firstname, lastname) VALUES (?, ?)`
const addTransactionQuery = `INSERT INTO transactions (event, amount, date, person) VALUES (?, ?, ?, ?)`

const selectAllTransactionsQuery = `SELECT * FROM transactions`
const selectAllPeopleQuery = `SELECT * FROM people`

const selectPersonQuery = `SELECT id FROM people WHERE firstname = ?`
const sumTransactionsQuery = `SELECT SUM(amount) FROM transactions WHERE person = ?`
