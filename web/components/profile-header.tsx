"use client";

export function ProfileHeader() {
  return (
    <header className="border-b border-border bg-card">
      <div className="mx-auto max-w-6xl px-4 py-12 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <div className="h-16 w-16 rounded-full bg-accent/10" />
            <div>
              <h1 className="text-3xl font-bold text-foreground">Alex Chen</h1>
              <p className="text-sm text-muted-foreground">0x742d...7e8f</p>
            </div>
          </div>
          <button className="rounded-lg bg-primary px-6 py-2 text-sm font-semibold text-primary-foreground transition-opacity hover:opacity-90">
            Edit Profile
          </button>
        </div>
      </div>
    </header>
  );
}
