import Image from "next/image";

export default function TopMenu() {
  return (
    <div className="w-full bg-white flex items-center   px-[80px]  ">
      <h1 className="text-[56px] font-bold text-emerald-700 drop-shadow-lg h-auto">
        Weather
      </h1>
      <Image
        src="/img/Icon.png"
        alt="logo"
        className="object-contain  drop-shadow-lg"
        width={200}
        height={200}
      />
    </div>
  );
}
