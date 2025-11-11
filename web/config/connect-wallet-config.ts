import { getDefaultConfig } from "@rainbow-me/rainbowkit";
import { hardhat } from "wagmi/chains";

const config = getDefaultConfig({
  appName: "My RainbowKit App",
  projectId: "ed2a5c1e433f2405cfae441c8f29bee9",
  chains: [hardhat],
  ssr: true, // If your dApp uses server side rendering (SSR)
});

export default config;
