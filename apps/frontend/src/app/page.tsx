import MainWeatherReport from "@/components/mainWeatherReport";
import Title from "@/components/title";
import dayjs, { Dayjs } from "dayjs";
// import Head from "next/head";

export default function Home() {
  let today = dayjs().format("YYYY-MM-DD");
  return (
    <div className="flex flex-col space-y-[48px] items-center justify-items-center bg-gradient-to-t p-8 pb-20 sm:p-20 font-[family-name:var(--font-geist-sans)] h-screen overflow-y-auto">
      {/* <div className="flex flex-col space-y-[48px] items-center justify-items-center bg-gradient-to-t p-8 pb-20 sm:p-20 font-[family-name:var(--font-geist-sans)] h-auto"> */}
      {/* <Head>
        <link rel="icon" href="/favicon.ico" sizes="any" />
      </Head> */}
      <Title />
      <MainWeatherReport />
      {/* <div className="text-white">{today}</div> */}
    </div>
  );
}
