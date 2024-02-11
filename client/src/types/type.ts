export interface Truck {
    id: number;
    name: string;
    userId: number;
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