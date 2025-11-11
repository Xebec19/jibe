import { RK_PROJECT_ID } from "@/lib/environments";
import { getDefaultConfig } from "@rainbow-me/rainbowkit";
import { hardhat } from "wagmi/chains";

const config = getDefaultConfig({
  appName: "jibe",
  projectId: RK_PROJECT_ID,
  chains: [hardhat],
  ssr: true, // If your dApp uses server side rendering (SSR)
});

export default config;
