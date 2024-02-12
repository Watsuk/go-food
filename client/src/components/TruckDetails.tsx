import { Button } from "@/components/ui/button";
import { Truck } from "@/types/type";
import ProductDetail from "./ProductDetail";

const testProduct = {
  id: 1,
  truckId: 1,
  name: "Burger",
  label: "Burger au boeuf avec frites et boisson",
  description:
    "Burger au boeuf avec frites et boisson, 100% fait maison, 100% délicieux, 100% satisfait ou remboursé.",
  price: 10,
  createdAt: "2022-01-01",
  updatedAt: "2022-01-01",
  deletedAt: "2022-01-01",
};

const TruckDetails = ({
  currentTruckData,
  onEdit,
}: {
  currentTruckData: Truck | null;
  onEdit: (truckId: number) => void;
}) => {
  if (!currentTruckData)
    return <div>Sélectionnez un camion pour voir les détails.</div>;

  return (
    <div className="flex flex-col gap-4 border border-gray-300 rounded-lg p-8">
      <div className="flex flex-row items-center gap-4"></div>
      <h1 className="text-2xl font-bold">{currentTruckData.name}</h1>
      <div className="flex flex-row items-center gap-4">
        <span>ouvert à : {currentTruckData.open_time}</span>
        <span>fermé à : {currentTruckData.close_time}</span>
      </div>
      <div className="flex flex-row items-center gap-4">
        <span>Créer le : {currentTruckData.created_at}</span>
      </div>
      <ProductDetail product={testProduct} />
      <span>Slots disponibles : {currentTruckData.slot_buffer}</span>
      <Button onClick={() => onEdit(currentTruckData.id)}>Commander</Button>
    </div>
  );
};

export default TruckDetails;
