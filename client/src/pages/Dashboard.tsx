import { getTrucks, getTrucksByTruckId } from "@/api/truck";
import { getUserById } from "@/api/user";
import HeaderPanel from "@/components/ShowUserInfo";
import TruckDetails from "@/components/TruckDetails";
import TruckList from "@/components/TruckList";
import TruckOwnerLayout from "@/components/TruckOwnerLayout";
import { ROLE_MAP, Truck, User } from "@/types/type"; // Assure-toi que le type User est correctement défini
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Dashboard() {
  const [selectedTruck, setSelectedTruck] = useState<Truck | null>(null);
  const [truckData, setTruckData] = useState<Truck[]>([]);
  const [userRole, setUserRole] = useState("");
  const [userData, setUserData] = useState<User | null>(null);

  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    const userId = localStorage.getItem("user_id");

    if (!token || !userId) {
      navigate("/auth/login");
      return;
    }

    const fetchUserData = async () => {
      try {
        const data = await getUserById(token, parseInt(userId));
        setUserData(data);
        const role = ROLE_MAP[data.role]; // Suppose que data.role est le numéro du rôle
        setUserRole(role); // Stocke le rôle en texte basé sur ROLE_MAP
      } catch (error) {
        console.error("Failed to fetch user data", error);
      }
    };

    fetchUserData();
  }, [navigate]);

  useEffect(() => {
    if (userRole !== "Restaurateur") {
      const fetchTrucks = async () => {
        const token = localStorage.getItem("token");
        if (token) {
          try {
            const trucks = await getTrucks(token);
            setTruckData(trucks);
          } catch (error) {
            console.error("Failed to fetch trucks", error);
          }
        }
      };

      fetchTrucks();
    }
  }, [userRole]);

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
      {userData?.username && <HeaderPanel userData={userData} />}
      {userRole === "Restaurateur" ? (
        <TruckOwnerLayout />
      ) : (
        <div className="h-full w-full p-8 flex flex-row items-center justify-between gap-4">
          {truckData ? (
            <>
              <TruckList
                onTruckSelect={handleTruckSelect}
                truckData={truckData}
              />
              <div className="w-3/4 h-full flex flex-col gap-4 border border-gray-300 rounded-lg p-8">
                <TruckDetails
                  currentTruckData={selectedTruck}
                  onEdit={() => console.log("Edit truck")}
                />
              </div>
            </>
          ) : (
            <div className="w-full text-center">Aucun camion trouvé.</div>
          )}
        </div>
      )}
    </div>
  );
}
