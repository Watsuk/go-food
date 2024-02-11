import TruckList from "@/components/TruckList";
import { Button } from "@/components/ui/button";
import { useEffect, useState } from "react";

export default function Sigin() {
  const [message, setMessage] = useState("");
  const [selectedTruck, setSelectedTruck] = useState(null);

  const handleTruckSelect = (truckId: number) => {};

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setMessage("You are not logged in");
      setTimeout(() => {
        window.location.href = "/auth/signin";
      }, 2000);
    }
    setMessage("");
  }, []);

  return (
    <div className="w-full h-full flex justify-center items-center gap-4 overflow-hidden">
      <p className="text-red-500">{message}</p>
      <div className="h-full w-full p-8 flex flex-row items-center justify-between gap-4">
        <TruckList
          currentTruck={selectedTruck}
          onTruckSelect={handleTruckSelect}
          truckData={[]}
        />
        <div className="w-3/4 h-full flex flex-col gap-4 border border-gray-300 rounded-lg p-8">
          <h1 className="text-2xl font-bold">Truck 7</h1>
          <div className="flex flex-col gap-4">
            <div className="flex flex-row items-center gap-4">
              <span>Open: 10:00</span>
              <span>Close: 20:00</span>
            </div>
            <div className="flex flex-row items-center gap-4">
              <span>Created: 2022-01-01</span>
              <span>Updated: 2022-01-01</span>
              <span>Deleted: 2022-01-01</span>
            </div>
            <Button>button</Button>
          </div>
        </div>
      </div>
    </div>
  );
}
