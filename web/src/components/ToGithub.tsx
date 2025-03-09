import React from "react"
import { GITHUB_URL, VERSION } from "@/lib/constant.ts"
import { Button } from "@/components/ui/button.tsx"
import { Github } from "lucide-react"

function ToGithub() {
  return (
    <Button
      onClick={() => open(GITHUB_URL)}
      variant="ghost"
      className="gap-1 underline underline-offset-4 has-[>svg]:px-2"
    >
      <Github />
      {VERSION}
    </Button>
  )
}

export default ToGithub
