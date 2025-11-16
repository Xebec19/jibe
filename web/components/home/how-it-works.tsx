'use client'

export function HowItWorks() {
  const steps = [
    {
      number: '1',
      title: 'Set Up',
      description: 'Connect your wallet and create your creator profile.',
    },
    {
      number: '2',
      title: 'Gate Content',
      description: 'Upload what you want to share and set token requirements.',
    },
    {
      number: '3',
      title: 'Share',
      description: 'Get a link to share with your community. They join with one click.',
    },
    {
      number: '4',
      title: 'Earn',
      description: 'Payments arrive instantly. Keep 100% of what you make.',
    },
  ]

  return (
    <section id="how-it-works" className="py-24 px-6 bg-muted/50">
      <div className="max-w-6xl mx-auto">
        <div className="text-center space-y-3 mb-16">
          <h2 className="text-4xl font-bold">How it works</h2>
          <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
            Launch your token-gated community in minutes.
          </p>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
          {steps.map((step, i) => (
            <div key={i}>
              <div className="space-y-3">
                <div className="text-4xl font-bold text-primary/40">
                  {step.number}
                </div>
                <h3 className="font-bold text-lg">{step.title}</h3>
                <p className="text-muted-foreground text-sm leading-relaxed">
                  {step.description}
                </p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}
