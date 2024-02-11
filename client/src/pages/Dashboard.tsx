import { getTrucks } from "@/api/truck";
import TruckDetails from "@/components/TruckDetails";
import TruckList from "@/components/TruckList";
import { Truck } from "@/types/type";
import { useEffect, useState } from "react";

const testTruckData: Truck[] = [
  {
    id: 1,
    name: "Truck 1",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 2,
    name: "Truck 2",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 3,
    name: "Truck 3",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 4,
    name: "Truck 4",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 5,
    name: "Truck 5",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 6,
    name: "Truck 6",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 7,
    name: "Truck 7",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 8,
    name: "Truck 8",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 9,
    name: "Truck 9",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 10,
    name: "Truck 10",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 11,
    name: "Truck 11",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 12,
    name: "Truck 12",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 13,
    name: "Truck 13",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
  {
    id: 14,
    name: "Truck 14",
    userId: 1,
    openTime: "10:00",
    closeTime: "20:00",
    createdAt: "2022-01-01",
    updatedAt: "2022-01-01",
    deletedAt: "2022-01-01",
    slotBuffer: 10,
  },
];

export default function Dashboard() {
  const [message, setMessage] = useState("");
  const [selectedTruck, setSelectedTruck] = useState<Truck | null>(null);
  const [truckData, setTruckData] = useState<Truck[]>([]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setMessage("You are not logged in");
      setTimeout(() => {
        window.location.href = "/auth/signin";
      }, 2000);
    }

    const fetchTrucks = async () => {
      try {
        const trucks = await getTrucks();
        setTruckData(trucks);
      } catch (error) {
        console.error("Failed to fetch trucks", error);
      }
    };

    fetchTrucks();
    setMessage("");
  }, []);

  const handleTruckSelect = (truckId: number) => {
    const truck = testTruckData.find((t) => t.id === truckId);
    setSelectedTruck(truck || null);

    /*
    const fetchCurrentTruck = async () => {
      try {
        const truck = await getTrucksByTruckId(truckId);
        setSelectedTruck(truck);
      } catch (error) {
        console.error("Failed to fetch truck", error);
      }
    };

    fetchCurrentTruck();

    if (!truck) {
      setMessage("Truck not found");
      return;
    }
    */
  };

  return (
    <div className="w-full h-full flex justify-center items-center gap-4 overflow-hidden">
      <p className="text-red-500">{message}</p>
      <div className="h-full w-full p-8 flex flex-row items-center justify-between gap-4">
        <TruckList
          onTruckSelect={handleTruckSelect}
          truckData={testTruckData}
        />
        <div className="w-3/4 h-full flex flex-col gap-4 border border-gray-300 rounded-lg p-8">
          <TruckDetails
            currentTruckData={selectedTruck}
            onEdit={() => console.log("Edit truck")}
          />
        </div>
      </div>
    </div>
  );
}
