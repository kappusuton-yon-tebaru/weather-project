import Image from "next/image";

export default function Bg() {
  return (
    <div className="bg-container">
      {/* Background Image */}
      <Image
        src="/img/bg2.jpg"
        alt="bg"
        fill={true}
        className="object-cover" // Ensures image covers the background
        priority={true} // Ensures the background loads quickly
      />

      {/* Gradient Overlay */}
      <div className="absolute inset-0 bg-black opacity-50"></div>
    </div>
  );
}
