<template>
  <div class="cashbunny-planner-wizard">
    <section
      class="cashbunny-planner-wizard__welcome"
      v-if="currentPage === PlannerWizardPages.Welcome"
    >
      <p>{{ t('cashbunny.planner.wizard.welcome.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__assets"
      v-else-if="currentPage === PlannerWizardPages.Assets"
    >
      <p>{{ t('cashbunny.planner.wizard.assets.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__liabilities"
      v-else-if="currentPage === PlannerWizardPages.Liabilities"
    >
      <p>{{ t('cashbunny.planner.wizard.liabilities.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__revenues"
      v-else-if="currentPage === PlannerWizardPages.Revenues"
    >
      <p>{{ t('cashbunny.planner.wizard.revenues.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__expenses"
      v-else-if="currentPage === PlannerWizardPages.Expenses"
    >
      <p>{{ t('cashbunny.planner.wizard.expenses.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__scheduled-transactions"
      v-else-if="currentPage === PlannerWizardPages.ScheduledTransactions"
    >
      <p>{{ t('cashbunny.planner.wizard.scheduledTransactions.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__savings-goals"
      v-else-if="currentPage === PlannerWizardPages.SavingsGoals"
    >
      <p>{{ t('cashbunny.planner.wizard.savingsGoals.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__notifications"
      v-else-if="currentPage === PlannerWizardPages.Notifications"
    >
      <p>{{ t('cashbunny.planner.wizard.notifications.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__suggestions"
      v-else-if="currentPage === PlannerWizardPages.Suggestions"
    >
      <p>{{ t('cashbunny.planner.wizard.suggestions.question') }}</p>
    </section>
    <section
      class="cashbunny-planner-wizard__complete"
      v-else-if="currentPage === PlannerWizardPages.Complete"
    >
      <p>{{ t('cashbunny.planner.wizard.complete.question') }}</p>
    </section>
    <div class="cashbunny-planner-wizard__nav-buttons">
      <ButtonComponent @click="onClickBack">
        <template v-if="currentPage === PlannerWizardPages.Welcome">
          {{ t('cashbunny.planner.wizard.welcome.no') }}
        </template>
        <template v-else>
          {{ t('common.back') }}
        </template>
      </ButtonComponent>
      <ButtonComponent type="success" @click="onClickContinue">
        <template v-if="currentPage === PlannerWizardPages.Welcome">
          {{ t('cashbunny.planner.wizard.welcome.yes') }}
        </template>
        <template v-else>
          {{ t('common.continue') }}
        </template>
      </ButtonComponent>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ButtonComponent from '@/core/components/ButtonComponent.vue'

enum PlannerWizardPages {
  Welcome,
  Assets,
  Liabilities,
  Revenues,
  Expenses,
  ScheduledTransactions,
  SavingsGoals,
  Notifications,
  Suggestions,
  Complete,
}

const { t } = useI18n()
const currentPage = ref<PlannerWizardPages>(PlannerWizardPages.Welcome)

const onClickBack = () => {
  if (currentPage.value !== 0) {
    currentPage.value -= 1
  }
}

const onClickContinue = () => {
  if (currentPage.value !== PlannerWizardPages.Complete) {
    currentPage.value += 1
  }
}
</script>

<style scoped lang="scss">
.cashbunny-planner-wizard {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;

  section {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;

    p {
      margin-top: 3em;
      white-space: pre-line;
      text-align: center;
      line-height: 1.7em;
      user-select: none;
      font-size: 1.2em;
    }
  }

  .cashbunny-planner-wizard__nav-buttons {
    display: flex;
    gap: 0.5em;
    margin-bottom: 3em;
    width: 75%;
    > button {
      flex: 1;
    }
  }

  section.cashbunny-planner-wizard__welcome {
    justify-content: center;
  }
}
</style>
