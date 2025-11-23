"use client";
import { useState, useEffect, useEffectEvent } from "react";
import "@rainbow-me/rainbowkit/styles.css";
import {
  AuthenticationStatus,
  createAuthenticationAdapter,
  getDefaultConfig,
  RainbowKitAuthenticationProvider,
  RainbowKitProvider,
} from "@rainbow-me/rainbowkit";
import { WagmiProvider, type Config } from "wagmi";
import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import { RK_PROJECT_ID } from "@/lib/environments";
import { hardhat } from "viem/chains";
import { createSiweMessage } from "viem/siwe";

type Props = {
  children: React.ReactNode;
};

function WagmiRBKitProvider({ children }: Props) {
  const [mounted, setMounted] = useState(false);
  const [config, setConfig] = useState<Config | null>(null);
  const [queryClient] = useState(() => new QueryClient());
  const [authenticationStatus, setAuthenticationStatus] =
    useState<AuthenticationStatus>("unauthenticated");

  const authenticationAdapter = createAuthenticationAdapter({
    getNonce: async () => {
      return "123asdfg";
    },

    createMessage: ({ nonce, address, chainId }) => {
      try {
        return createSiweMessage({
          domain: window.location.host,
          address,
          statement: "Sign in with Ethereum in the app.",
          uri: window.location.origin,
          version: "1",
          chainId,
          nonce,
        });
      } catch (err) {
        console.error(err);
      }
    },

    verify: async ({ message, signature }) => {
      setAuthenticationStatus("authenticated");
      // todo hit and api here to verify message and signature in backend
      return true;
    },

    signOut: async () => {
      setAuthenticationStatus("unauthenticated");
      // todo hit an api here to signou
    },
  });

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
        <RainbowKitAuthenticationProvider
          adapter={authenticationAdapter}
          status={authenticationStatus}
        >
          <RainbowKitProvider>{children}</RainbowKitProvider>
        </RainbowKitAuthenticationProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}

export default WagmiRBKitProvider;
