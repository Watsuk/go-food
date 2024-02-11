export interface Truck {
    id: number;
    name: string;
    userId: number;
    slotBuffer: number;
    openTime: string;
    closeTime: string;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
}

export interface User {
    id: number;
    username: string;
    email: string;
    password: string;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
}

export interface Product {
    id: number;
    truckId: Truck["id"];
    name: string;
    label: string;
    description: string;
    price: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
}