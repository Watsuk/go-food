const Navbar = () => {
  return (
    <nav className="w-full bg-black py-4 px-6 flex justify-between items-center">
      <div className="text-white font-bold text-lg">go food</div>
      <button
        className="text-white"
        onClick={() => {
          localStorage.removeItem("token");
          window.location.href = "/auth/signin";
        }}
      >
        Logout
      </button>
    </nav>
  );
};

export default Navbar;
