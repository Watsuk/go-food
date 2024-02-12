import React, { useEffect, useState } from "react";
import { getProductByTruckId, createProduct } from "@/api/truck"; // Vérifiez les chemins d'importation
import { ProductData, Truck } from "@/types/type"; // Vérifiez les chemins d'importation
import CreateProductForm from "@/components/CreateProductForm"; // Supposons l'existence de ce formulaire

const TruckMenu = ({ truck }: { truck: Truck }) => {
  const [products, setProducts] = useState<ProductData[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const jwt = localStorage.getItem("token");
    if (!jwt) {
      console.error("JWT missing");
      return;
    }

    const fetchProducts = async () => {
      try {
        const productsData = await getProductByTruckId(jwt, truck.id);
        setProducts(productsData);
      } catch (error) {
        console.error("Failed to fetch products:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchProducts();
  }, [truck.id]);

  if (loading) {
    return <div>Chargement du menu...</div>;
  }

  return (
    <div className="menu-container">
      {products.length > 0 ? (
        <div>
          <h2>Menu du Camion {truck.name}</h2>
          {products.map((product, index) => (
            <div key={index}>
              <h3>{product.label}</h3>
              <p>{product.description}</p>
              <p>Prix: {product.price}</p>
            </div>
          ))}
        </div>
      ) : (
        <div>
          <h2>Aucun produit trouvé pour {truck.name}</h2>
          <CreateProductForm truckID={truck.id} />
        </div>
      )}
    </div>
  );
};

export default TruckMenu;
