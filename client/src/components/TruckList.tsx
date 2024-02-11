import { Truck } from "@/types/type";
import { ScrollArea } from "@/components/ui/scroll-area";

function TruckList({
  currentTruck,
  onTruckSelect,
  truckData,
}: {
  currentTruck: Truck | null;
  onTruckSelect: (truckId: number) => void;
  truckData: Truck[];
}) {
  return (
    <ScrollArea className="h-full w-1/4 border border-gray-300 rounded-lg overflow-hidden">
      {truckData.length === 0 && (
        <div className="flex items-center justify-center h-full">
          No trucks available
        </div>
      )}
      {truckData.map((truck) => (
        <div
          key={truck.id}
          onClick={() => onTruckSelect(truck.id)}
          className="flex flex-col items-start justify-center gap-4 p-4 border border-gray-300 w-full hover:bg-gray-100 cursor-pointer"
        >
          <h3>{truck.name}</h3>
          <span className="flex gap-4">
            <span>Open: {truck.openTime}</span>
            <span>Close: {truck.closeTime}</span>
          </span>
        </div>
      ))}
    </ScrollArea>
  );
}

export default TruckList;
