import { ProfileHeader } from "@/components/profile-header";
import { BalanceCard } from "@/components/balance-card";
import { NFTGrid } from "@/components/nft-grid";
import { TransactionHistory } from "@/components/transaction-history";

export default function ProfilePage() {
  return (
    <main className="min-h-screen bg-background">
      <ProfileHeader />

      <div className="mx-auto max-w-6xl px-4 py-16 sm:px-6 lg:px-8">
        {/* Balance Section */}
        <div className="mb-20">
          <BalanceCard />
        </div>

        {/* NFTs and Transactions */}
        <div className="grid gap-20 lg:grid-cols-3">
          <div className="lg:col-span-2">
            <NFTGrid />
          </div>
          <div>
            <TransactionHistory />
          </div>
        </div>
      </div>
    </main>
  );
}
