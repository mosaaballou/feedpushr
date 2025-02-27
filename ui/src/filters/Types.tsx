
export interface FilterProps {
    nbSuccess?: number
    nbError?: number
    [key: string]: any
}

export interface Filter {
  id: number
  alias: string
  name: string
  desc: string
  enabled: boolean
  tags: string[]
  props: FilterProps
}

export interface PropsSpec {
  name: string
  desc: string
  type: string
}

export interface FilterSpec {
  name: string
  desc: string
  props: PropsSpec[]
}

export interface FilterForm {
  id?: number
  alias: string
  name: string
  tags: string[]
  props: FilterProps
}
