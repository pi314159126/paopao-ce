import{_ as q}from"./whisper-6e97e8e3.js";import{_ as D,a as R}from"./post-item.vue_vue_type_style_index_0_lang-299dc5b6.js";import{_ as U}from"./post-skeleton-dd55abe8.js";import{_ as E}from"./main-nav.vue_vue_type_style_index_0_lang-52234344.js";import{u as G}from"./vuex-44de225f.js";import{b as J}from"./vue-router-e5a2430e.js";import{W as L}from"./v3-infinite-loading-2c58ec2f.js";import{T as Y,u as K,f as Q,_ as X}from"./index-7d4ab953.js";import{d as Z,H as t,b as ee,f as n,k as a,w as u,q as d,Y as h,e as o,bf as f,F as S,u as $,j as z,x as oe}from"./@vue-a481fc63.js";import{F as se,G as te,a as ne,J as ae,k as ie,H as le}from"./naive-ui-eecf2ec3.js";import"./content-aad6a645.js";import"./@vicons-f0266f88.js";import"./paopao-video-player-2fe58954.js";import"./copy-to-clipboard-4ef7d3eb.js";import"./@babel-725317a4.js";import"./toggle-selection-93f4ad84.js";import"./vooks-6d99783e.js";import"./evtd-b614532e.js";import"./axios-4a70c6fc.js";import"./moment-2ab8298d.js";/* empty css               */import"./seemly-76b7b838.js";import"./vueuc-7c8d4b48.js";import"./@css-render-7124a1a5.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";const re={key:0,class:"skeleton-wrap"},_e={key:1},ue={key:0,class:"empty-wrap"},ce={key:1},pe={key:2},me={class:"load-more-wrap"},de={class:"load-more-spinner"},fe=Z({__name:"Collection",setup(ve){const v=G(),A=J(),B=se(),c=t(!1),_=t(!1),s=t([]),l=t(+A.query.p||1),w=t(20),p=t(0),g=t(!1),k=t({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1}),y=e=>{k.value=e,g.value=!0},I=()=>{g.value=!1},x=e=>{B.success({title:"提示",content:"确定"+(e.user.is_following?"取消关注":"关注")+"该用户吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{e.user.is_following?K({user_id:e.user.id}).then(r=>{window.$message.success("操作成功"),C(e.user_id,!1)}).catch(r=>{}):Q({user_id:e.user.id}).then(r=>{window.$message.success("关注成功"),C(e.user_id,!0)}).catch(r=>{})}})};function C(e,r){for(let m in s.value)s.value[m].user_id==e&&(s.value[m].user.is_following=r)}const b=()=>{c.value=!0,Y({page:l.value,page_size:w.value}).then(e=>{c.value=!1,e.list.length===0&&(_.value=!0),l.value>1?s.value=s.value.concat(e.list):(s.value=e.list,window.scrollTo(0,0)),p.value=Math.ceil(e.pager.total_rows/w.value)}).catch(e=>{c.value=!1,l.value>1&&l.value--})},M=()=>{l.value<p.value||p.value==0?(_.value=!1,l.value++,b()):_.value=!0};return ee(()=>{b()}),(e,r)=>{const m=E,O=U,P=ae,T=D,F=le,H=R,N=q,V=te,W=ie,j=ne;return o(),n("div",null,[a(m,{title:"收藏"}),a(V,{class:"main-content-wrap",bordered:""},{default:u(()=>[c.value&&s.value.length===0?(o(),n("div",re,[a(O,{num:w.value},null,8,["num"])])):(o(),n("div",_e,[s.value.length===0?(o(),n("div",ue,[a(P,{size:"large",description:"暂无数据"})])):h("",!0),f(v).state.desktopModelShow?(o(),n("div",ce,[(o(!0),n(S,null,$(s.value,i=>(o(),d(F,{key:i.id},{default:u(()=>[a(T,{post:i,isOwner:f(v).state.userInfo.id==i.user_id,addFollowAction:!0,onSendWhisper:y,onHandleFollowAction:x},null,8,["post","isOwner"])]),_:2},1024))),128))])):(o(),n("div",pe,[(o(!0),n(S,null,$(s.value,i=>(o(),d(F,{key:i.id},{default:u(()=>[a(H,{post:i,isOwner:f(v).state.userInfo.id==i.user_id,addFollowAction:!0,onSendWhisper:y,onHandleFollowAction:x},null,8,["post","isOwner"])]),_:2},1024))),128))]))])),a(N,{show:g.value,user:k.value,onSuccess:I},null,8,["show","user"])]),_:1}),p.value>0?(o(),d(j,{key:0,justify:"center"},{default:u(()=>[a(f(L),{class:"load-more",slots:{complete:"没有更多收藏了",error:"加载出错"},onInfinite:M},{spinner:u(()=>[z("div",me,[_.value?h("",!0):(o(),d(W,{key:0,size:14})),z("span",de,oe(_.value?"没有更多收藏了":"加载更多"),1)])]),_:1})]),_:1})):h("",!0)])}}});const Ye=X(fe,[["__scopeId","data-v-735372fb"]]);export{Ye as default};
