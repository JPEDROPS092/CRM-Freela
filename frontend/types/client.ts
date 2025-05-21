export interface Client {
  id: number
  name: string
  email: string
  phone: string
  company: string
  notes?: string
  status: string
  created_at: string
  updated_at?: string
}

export interface ClientsResponse {
  clients: Client[]
  total: number
  totalPages: number
  currentPage: number
}

export interface ClientStats {
  total: number
  active: number
  inactive: number
  archived: number
}
