<template>
  <div class="faq-section">
    <div class="faq-container">
      <div v-for="(category, index) in faqCategories" :key="index" class="faq-category">
        <h3>{{ category.title }}</h3>
        
        <div class="faq-items">
          <div 
            v-for="(item, itemIndex) in category.items" 
            :key="itemIndex" 
            class="faq-item"
            :class="{ 'active': openItems[`${index}-${itemIndex}`] }"
          >
            <div class="faq-question" @click="toggleItem(index, itemIndex)">
              <span>{{ item.question }}</span>
              <div class="icon-container">
                <i class="ri-add-line" v-if="!openItems[`${index}-${itemIndex}`]"></i>
                <i class="ri-subtract-line" v-else></i>
              </div>
            </div>
            
            <div v-if="openItems[`${index}-${itemIndex}`]" class="faq-answer">
              <div v-html="item.answer"></div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="faq-footer">
        <div class="faq-footer-content">
          <div class="icon">
            <i class="ri-question-line"></i>
          </div>
          <div class="content">
            <p>Hittar du inte svar på din fråga?</p>
            <a href="#contact-form" class="btn btn-primary">Kontakta oss</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'FAQSection',
  setup() {
    const openItems = ref({});
    
    const toggleItem = (categoryIndex, itemIndex) => {
      const key = `${categoryIndex}-${itemIndex}`;
      openItems.value[key] = !openItems.value[key];
    };
    
    const faqCategories = [
      {
        title: 'Allmänna frågor',
        items: [
          {
            question: 'Vad är Svensk Hälsovård?',
            answer: 'Svensk Hälsovård är en ledande leverantör av hälsokontroller och blodprovsanalyser i Sverige. Vi erbjuder både engångsundersökningar och prenumerationstjänster för kontinuerlig uppföljning av din hälsa. Vårt erfarna team av läkare och sjuksköterskor garanterar högkvalitativ service och professionell rådgivning.'
          },
          {
            question: 'Hur fungerar processen?',
            answer: 'Processen är enkel! Välj den tjänst du är intresserad av och genomför köpet. Du kommer sedan att få en bekräftelse via e-post med information om din bokning. För blodprov och hälsokontroller kommer du att kontaktas för att boka en tid. Efter genomförd undersökning kommer du att få tillgång till dina resultat digitalt, samt en personlig uppföljning av en av våra experter.'
          },
          {
            question: 'Är era tjänster tillgängliga i hela landet?',
            answer: 'Ja, våra tjänster finns tillgängliga i hela Sverige. Vi har partnerlaboratorier och kliniker på över 30 orter i landet. Om du inte har möjlighet att besöka en av våra kliniker kan vi i vissa fall även erbjuda hembesök för provtagning (gäller endast utvalda områden och tjänster).'
          }
        ]
      },
      {
        title: 'Tjänster och priser',
        items: [
          {
            question: 'Vad kostar era hälsokontroller?',
            answer: 'Priserna för våra hälsokontroller varierar beroende på omfattning och innehåll. Vår baskontroll börjar från 995 kr och våra mer omfattande hälsokontroller kostar från 2995 kr. Du hittar detaljerad prisinformation på respektive tjänstsida. Vi erbjuder även företagspaket med specialpriser.'
          },
          {
            question: 'Vilka blodprover ingår i era olika tjänster?',
            answer: 'Våra tjänster inkluderar olika blodprovsanalyser beroende på vilken tjänst du väljer. Baspaket inkluderar vanligtvis hemoglobin, blodsocker, kolesterol samt lever- och njurvärden. Våra mer omfattande paket kan inkludera hormoner, vitaminer, mineraler, inflammationsmarkörer och mycket mer. Du kan se exakt vilka prover som ingår i respektive tjänst på dess detaljsida.'
          },
          {
            question: 'Hur fungerar prenumerationstjänsterna?',
            answer: 'Våra prenumerationstjänster ger dig möjlighet att regelbundet kontrollera din hälsa över tid. När du väljer en prenumeration kommer du att få genomföra samma undersökning med jämna intervall (normalt var 3:e, 6:e eller 12:e månad). Du betalar ett förmånligt prenumerationspris och kan när som helst avsluta din prenumeration. Efter varje undersökning får du tillgång till dina resultat samt en jämförelse med dina tidigare resultat för att se din hälsoutveckling över tid.'
          }
        ]
      },
      {
        title: 'Betalning och bokning',
        items: [
          {
            question: 'Vilka betalningsmetoder accepterar ni?',
            answer: 'Vi accepterar betalning via Svea Ekonomi, som erbjuder säkra betalningslösningar. Du kan betala med bankkort (Visa/Mastercard) eller faktura. För prenumerationstjänster sker betalningen automatiskt inför varje ny period.'
          },
          {
            question: 'Hur avbokar eller ändrar jag min bokade tid?',
            answer: 'Du kan avboka eller ändra din bokade tid via e-post eller telefon senast 24 timmar innan din bokade tid utan kostnad. För avbokningar med kortare varsel debiteras en avgift på 50% av tjänstens pris. Kontakta vår kundtjänst för att avboka eller ändra din tid.'
          },
          {
            question: 'Hur snabbt kan jag få en tid för min undersökning?',
            answer: 'Tillgängliga tider varierar beroende på säsong och efterfrågan. Normalt kan vi erbjuda tider inom 1-2 veckor från bokningstillfället. I akuta fall kan vi ofta erbjuda tider med kortare varsel. Efter genomfört köp kommer du att kontaktas för att boka en tid som passar dig.'
          }
        ]
      },
      {
        title: 'Resultat och uppföljning',
        items: [
          {
            question: 'Hur lång tid tar det innan jag får mina resultat?',
            answer: 'Resultaten från blodprover är normalt tillgängliga inom 2-3 arbetsdagar. För mer omfattande hälsokontroller kan det ta upp till 5 arbetsdagar innan alla resultat är sammanställda. Du får en notifikation via e-post när dina resultat finns tillgängliga i din personliga hälsoportal.'
          },
          {
            question: 'Hur får jag tillgång till mina resultat?',
            answer: 'Du får tillgång till dina resultat genom vår säkra digitala hälsoportal. Du loggar in med BankID för att säkerställa att endast du har tillgång till dina personliga hälsodata. Här kan du se dina aktuella resultat, jämföra med tidigare resultat och se trender över tid. Du kan även ladda ner dina resultat som PDF för att dela med din läkare om du önskar.'
          },
          {
            question: 'Vad händer om något avvikande upptäcks i mina prover?',
            answer: 'Om något avvikande upptäcks i dina prover kommer du att kontaktas personligen av en av våra läkare eller sjuksköterskor. De går igenom resultaten med dig och ger rekommendationer om eventuella ytterligare åtgärder. Om det behövs kan vi hjälpa till med remiss till specialistvård eller boka uppföljande prover.'
          }
        ]
      },
    ];
    
    return {
      faqCategories,
      openItems,
      toggleItem
    };
  }
};
</script>

<style lang="scss" scoped>
.faq-section {
  .faq-container {
    max-width: 900px;
    margin: 0 auto;
  }
  
  .faq-category {
    margin-bottom: 2.5rem;
    
    h3 {
      margin-bottom: 1.5rem;
      color: $primary;
      font-size: 1.5rem;
      font-weight: 700;
      position: relative;
      padding-left: 1.5rem;
      
      &::before {
        content: '';
        position: absolute;
        left: 0;
        top: 0.5rem;
        width: 5px;
        height: 70%;
        background-color: $primary;
        border-radius: 5px;
      }
    }
    
    .faq-items {
      .faq-item {
        margin-bottom: 1rem;
        background-color: $light;
        border-radius: $border-radius-lg;
        box-shadow: $box-shadow;
        overflow: hidden;
        transition: $transition-base;
        
        &:hover {
          transform: translateY(-3px);
        }
        
        &.active {
          box-shadow: $box-shadow-hover;
          
          .faq-question {
            background-color: rgba($primary, 0.05);
            color: $primary;
            
            .icon-container {
              background-color: $primary;
              color: $light;
            }
          }
        }
        
        .faq-question {
          padding: 1.5rem;
          display: flex;
          justify-content: space-between;
          align-items: center;
          cursor: pointer;
          font-weight: 600;
          transition: $transition-base;
          
          &:hover {
            background-color: rgba($primary, 0.05);
          }
          
          .icon-container {
            display: flex;
            align-items: center;
            justify-content: center;
            width: 28px;
            height: 28px;
            border-radius: 50%;
            background-color: #e2e8f0;
            color: #64748b;
            transition: $transition-base;
            flex-shrink: 0;
            
            i {
              font-size: 1rem;
            }
          }
        }
        
        .faq-answer {
          padding: 0 1.5rem 1.5rem;
          color: #4a5568;
          line-height: 1.7;
          
          p {
            margin-bottom: 1rem;
            
            &:last-child {
              margin-bottom: 0;
            }
          }
          
          ul {
            padding-left: 1.5rem;
            margin-bottom: 1rem;
            
            li {
              margin-bottom: 0.5rem;
            }
          }
        }
      }
    }
  }
  
  .faq-footer {
    margin-top: 3rem;
    
    .faq-footer-content {
      background-color: $light;
      border-radius: $border-radius-lg;
      box-shadow: $box-shadow;
      padding: 2rem;
      display: flex;
      align-items: center;
      gap: 2rem;
      
      @media (max-width: $breakpoint-sm) {
        flex-direction: column;
        text-align: center;
      }
      
      .icon {
        width: 70px;
        height: 70px;
        background-color: $primary-lighter;
        color: $primary;
        font-size: 2rem;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-shrink: 0;
      }
      
      .content {
        flex: 1;
        
        p {
          margin-bottom: 1rem;
          font-size: 1.1rem;
          font-weight: 500;
          color: $dark;
        }
        
        .btn {
          display: inline-flex;
          padding: 0.75rem 1.5rem;
          font-weight: 600;
          box-shadow: $box-shadow;
          align-items: center;
          gap: 0.5rem;
          
          &:hover {
            transform: translateY(-3px);
            box-shadow: $box-shadow-hover;
          }
        }
      }
    }
  }
}
</style>