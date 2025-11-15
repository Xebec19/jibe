"use client";
import { useState, useEffect, useEffectEvent } from "react";
import "@rainbow-me/rainbowkit/styles.css";
import { getDefaultConfig, RainbowKitProvider } from "@rainbow-me/rainbowkit";
import { WagmiProvider, type Config } from "wagmi";
import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import { RK_PROJECT_ID } from "@/lib/environments";
import { hardhat } from "viem/chains";

type Props = {
  children: React.ReactNode;
};

function WagmiRBKitProvider({ children }: Props) {
  const [mounted, setMounted] = useState(false);
  const [config, setConfig] = useState<Config | null>(null);
  const [queryClient] = useState(() => new QueryClient());

  const setupConfig = useEffectEvent(() => {
    // Only create config on client-side after mount
    if (typeof window !== "undefined") {
      const wagmiConfig = getDefaultConfig({
        appName: "jibe",
        projectId: RK_PROJECT_ID,
        chains: [hardhat],
        ssr: true,
      });
      setConfig(wagmiConfig);
      setMounted(true);
    }
  });

  useEffect(() => {
    setupConfig();
  }, []);

  if (!mounted || !config) {
    return null;
  }

  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <RainbowKitProvider>{children}</RainbowKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}

export default WagmiRBKitProvider;
