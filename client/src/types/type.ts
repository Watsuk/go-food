export interface Truck {
    id: number;
    name: string;
    user_id: number;
    slot_buffer: number;
    open_time: string;
    close_time: string;
    created_at: string;
    updated_at: string;
    deleted_at: string;
}

export interface User {
    id: number;
    username: string;
    email: string;
    role: number;
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