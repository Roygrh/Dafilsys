package models

type Customer struct {
    ID string
    Name string
    Email string
    PaymentMethods []PaymentMethod
}