import { Product } from "@/types/type";

const ProductDetail = ({ product }: { product: Product }) => {
  return (
    <div className="flex flex-col gap-2 p-4 border border-gray-200 rounded-lg">
      <h1 className="text-2xl font-bold">{product.name}</h1>
      <h3 className="text-lg font-semibold">{product.label}</h3>
      <p>{product.description}</p>
      <div className="flex justify-between items-center">
        <span>Prix : {product.price}</span>
        {/*<span>Quantité : {product.quantity}</span>*/}
      </div>
    </div>
  );
};

export default ProductDetail;
