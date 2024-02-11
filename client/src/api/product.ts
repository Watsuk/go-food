const API_URL = "http://localhost:3000";

interface ProductData {
    truckID: number;
    label: string;
    description: string;
    price: number;
  }

export const createProduct = async (productData: ProductData): Promise<ProductData> => {
    const response = await fetch(`${API_URL}/product`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        truck_id: productData.truckID, 
        label: productData.label,
        description: productData.description,
        price: productData.price,
      }),
    });
  
    if (!response.ok) {
      throw new Error("Failed to create product");
    }
  
    return response.json(); 
  };

interface Product {
    id: number;
    truckID: number;
    label: string;
    description: string;
    price: number;
    createdAt: string;
    updatedAt: string;
}

export const getProductById = async (productID: number): Promise<Product> => {
    const response = await fetch(`${API_URL}/product/${productID}`, {
      method: "GET",
      headers: {
        "Accept": "application/json",
      },
    });
  
    if (!response.ok) {
      throw new Error("Failed to fetch product");
    }
  
    return response.json();
  };
  
  
  