import { createRouter, createWebHistory } from 'vue-router'
import CreateBankView from '../views/CreateBankView.vue'
import JoinBankView from '../views/JoinBankView.vue'
import AllBankView from '../views/AllBankView.vue'
import MyBankView from '../views/MyBankView.vue'
import PersonalBankView from '../views/PersonalBankView.vue'
import TakeLoanCustomerView from '../views/TakeLoanCustomerView.vue'
import MyLoanCustomerView from '../views/MyLoanCustomerView.vue'
import MyLoanBankView from '../views/MyLoanBankView'
import LoanDetailsView from '../views/LoanDetailsView.vue'
import BankLoanDetailsView from '../views/BankLoanDetailsView.vue'
import AuctionListView from '../views/AuctionListView.vue'
import AuctionItemView from '../views/AuctionItemView.vue'

const routes = [
  {
    path: '/create_bank_view',
    name: 'CreateBankView',
    component: CreateBankView,
  },
  {
    path: '/join_bank_view',
    name: 'JoinBankView',
    component: JoinBankView,
  },
  {
    path: '/all_bank_view',
    name: 'AllBankView',
    component: AllBankView,
  },
  {
    path: '/my_bank_view',
    name: 'MyBankView',
    component: MyBankView,
  },
  {
    path: '/personal_bank_view/:bankId',
    name: 'PersonalBankView',
    component: PersonalBankView,
  },
  {
    path: '/take_loan_customer_view',
    name: 'TakeLoanCustomerView',
    component: TakeLoanCustomerView,
  },
  {
    path: '/my_loan_customer_view',
    name: 'MyLoanCustomerView',
    component: MyLoanCustomerView,
  },
  {
    path: '/my_loan_bank_view',
    name: 'MyLoanBankView',
    component: MyLoanBankView,
  },
  {
    path: '/loan_details_view/:bankId/:loanId/:customerAddress',
    name: 'LoanDetailsView',
    component: LoanDetailsView,
  },
  {
    path: '/bank_loan_details_View/:bankId/:loanId/:customerAddress',
    name: 'BankLoanDetailsView',
    component: BankLoanDetailsView,
  },
  {
    path: '/auction_list_view',
    name: 'AuctionListView',
    component: AuctionListView
  },
  {
    path: '/auction_item_view/:loanId',
    name: 'AuctionItemView',
    component: AuctionItemView,
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})


export default router
