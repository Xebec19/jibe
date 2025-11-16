'use client'

import { Button } from '@/components/ui/button'
import { ArrowRight } from 'lucide-react'

export function CTA() {
  return (
    <section className="py-24 px-6">
      <div className="max-w-3xl mx-auto text-center space-y-8">
        <div className="space-y-4">
          <h2 className="text-4xl font-bold text-balance">
            Ready to own your community?
          </h2>
          <p className="text-lg text-muted-foreground max-w-2xl mx-auto text-balance">
            Join hundreds of creators earning with their most engaged followers.
          </p>
        </div>

        <div className="flex flex-col sm:flex-row gap-3 justify-center">
          <Button className="bg-primary text-primary-foreground hover:bg-primary/90 px-6 h-11 text-base font-medium">
            Create Your Space
            <ArrowRight className="ml-2" size={18} />
          </Button>
          <Button
            variant="outline"
            className="border-foreground/20 px-6 h-11 text-base font-medium hover:bg-muted"
          >
            Learn More
          </Button>
        </div>
      </div>
    </section>
  )
}
