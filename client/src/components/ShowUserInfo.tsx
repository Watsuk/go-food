import { User } from "@/types/type";
import React, { useEffect, useState } from "react";

interface HeaderPanelProps {
  userData: User | null;
}

const HeaderPanel: React.FC<HeaderPanelProps> = ({ userData }) => {
  const [editing, setEditing] = useState(false);
  const [user, setUser] = useState(userData);

  useEffect(() => {
    setUser(userData);
  }, [userData]);

  const handleEdit = () => {
    setEditing(true);
  };

  const handleSave = () => {
    // Enregistrer les modifications ici
    setEditing(false);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (user) {
      setUser({ ...user, [e.target.name]: e.target.value });
    }
  };
  return (
    <div className="border-b p-4 flex flex-row justify-between w-full">
      <div className="flex items-center justify-between">
        <h1 className="text-xl font-bold mr-2">User Info</h1>
        {!editing && (
          <button
            className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded"
            onClick={handleEdit}
          >
            Edit
          </button>
        )}
      </div>
      <div className="mt-4">
        <label className="block mb-2 font-bold">Name:</label>
        {editing ? (
          <input
            type="text"
            name="name"
            value={user?.username}
            onChange={handleChange}
            className="border border-gray-400 p-2 rounded"
          />
        ) : (
          <p>{user?.username}</p>
        )}
      </div>
      <div className="mt-4">
        <label className="block mb-2 font-bold">Role:</label>
        {editing ? (
          <input
            type="text"
            name="role"
            value={user?.role}
            onChange={handleChange}
            className="border border-gray-400 p-2 rounded"
          />
        ) : (
          <p>{user?.role}</p>
        )}
      </div>
      <div className="mt-4">
        <label className="block mb-2 font-bold">Email:</label>
        {editing ? (
          <input
            type="email"
            name="email"
            value={user?.email}
            onChange={handleChange}
            className="border border-gray-400 p-2 rounded"
          />
        ) : (
          <p>{user?.email}</p>
        )}
      </div>
      <div className="mt-4">
        <label className="block mb-2 font-bold">Password:</label>
        {editing ? (
          <input
            type="password"
            name="password"
            value={user?.role}
            onChange={handleChange}
            className="border border-gray-400 p-2 rounded"
          />
        ) : (
          <p>{user?.role}</p>
        )}
      </div>
      {editing && (
        <button
          className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mt-4"
          onClick={handleSave}
        >
          Save
        </button>
      )}
    </div>
  );
};

export default HeaderPanel;
