import { getTrucks, getTrucksByTruckId } from "@/api/truck";
import { getUserById } from "@/api/user";
import HeaderPanel from "@/components/ShowUserInfo";
import TruckDetails from "@/components/TruckDetails";
import TruckList from "@/components/TruckList";
import { Truck, User } from "@/types/type"; // Assure-toi que le type User est correctement défini
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Dashboard() {
  const [selectedTruck, setSelectedTruck] = useState<Truck | null>(null);
  const [truckData, setTruckData] = useState<Truck[]>([]);
  const [userData, setUserData] = useState<User | null>(null);

  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    const userId = localStorage.getItem("user_id"); // Assure-toi que l'ID de l'utilisateur est stocké sous le nom "userId"

    if (!token || !userId) {
      navigate("/auth/login");
      return; // Stoppe l'exécution si le token ou l'userId est manquant
    }

    const fetchUserData = async () => {
      try {
        const data = await getUserById(token!, parseInt(userId!));
        setUserData(data);
        console.error(data);
      } catch (error) {
        console.error("Failed to fetch user data", error);
      }
    };

    const fetchTrucks = async () => {
      try {
        const trucks = await getTrucks(token!);
        setTruckData(trucks);
      } catch (error) {
        console.error("Failed to fetch trucks", error);
      }
    };

    fetchUserData();
    fetchTrucks();
  }, [navigate]);

  const handleTruckSelect = async (truckId: number) => {
    const token = localStorage.getItem("token");

    if (truckId === selectedTruck?.id) {
      return;
    }

    try {
      const truck = await getTrucksByTruckId(token!, truckId);
      setSelectedTruck(truck);
    } catch (error) {
      console.error("Failed to fetch truck", error);
    }
  };

  return (
    <div className="w-full h-full flex flex-col justify-center items-center gap-4 overflow-hidden">
      {userData?.username && <HeaderPanel userData={userData} />}{" "}
      <div className="h-full w-full p-8 flex flex-row items-center justify-between gap-4">
        <TruckList onTruckSelect={handleTruckSelect} truckData={truckData} />
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
