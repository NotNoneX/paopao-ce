import{p as B,a as N,_ as V,b as D,c as I}from"./content-555eec37.js";import{d as F,n as P,a3 as A,o as i,c as k,V as o,a7 as E,a1 as s,a as u,F as Q,a4 as R,a2 as f,_ as p,e as _,M as m,Q as r,O as c,s as S}from"./@vue-e0e89260.js";import{u as G}from"./vuex-473b3783.js";import{u as J}from"./vue-router-b8e3382f.js";import{c as K}from"./formatTime-4210fcd1.js";import{a as oe}from"./copy-to-clipboard-1dd3075d.js";import{i as ie,j as U,l as W,m as X,o as le}from"./@vicons-0524c43e.js";import{j as y,o as Y,O as Z,e as ue,P as re,a as ee,M as te}from"./naive-ui-62663ad7.js";const ce={class:"post-item"},pe={class:"nickname-wrap"},_e={class:"username-wrap"},me={class:"timestamp-mobile"},de={class:"item-header-extra"},ve=["innerHTML"],he={class:"opt-item"},ge={class:"opt-item"},He=F({__name:"mobile-post-item",props:{post:{}},setup(C){const q=C,h=J(),T=G(),t=l=>()=>S(y,null,{default:()=>S(l)}),x=P(()=>[{label:"复制链接",key:"copyTweetLink",icon:t(le)}]),O=async l=>{switch(l){case"copyTweetLink":oe(`${window.location.origin}/#/post?id=${e.value.id}`),window.$message.success("链接已复制到剪贴板");break}},e=P(()=>{let l=Object.assign({texts:[],imgs:[],videos:[],links:[],attachments:[],charge_attachments:[]},q.post);return l.contents.map(n=>{(+n.type==1||+n.type==2)&&l.texts.push(n),+n.type==3&&l.imgs.push(n),+n.type==4&&l.videos.push(n),+n.type==6&&l.links.push(n),+n.type==7&&l.attachments.push(n),+n.type==8&&l.charge_attachments.push(n)}),l}),a=l=>{h.push({name:"post",query:{id:l}})},v=(l,n)=>{if(l.target.dataset.detail){const d=l.target.dataset.detail.split(":");if(d.length===2){T.commit("refresh"),d[0]==="tag"?h.push({name:"home",query:{q:d[1],t:"tag"}}):h.push({name:"user",query:{s:d[1]}});return}}a(n)};return(l,n)=>{const d=Y,M=A("router-link"),w=Z,L=ue,$=re,b=N,j=V,g=D,H=I,se=ee,ae=te;return i(),k("div",ce,[o(ae,{"content-indented":""},E({avatar:s(()=>[o(d,{round:"",size:30,src:e.value.user.avatar},null,8,["src"])]),header:s(()=>[u("span",pe,[o(M,{onClick:n[0]||(n[0]=f(()=>{},["stop"])),class:"username-link",to:{name:"user",query:{s:e.value.user.username}}},{default:s(()=>[_(m(e.value.user.nickname),1)]),_:1},8,["to"])]),u("span",_e," @"+m(e.value.user.username),1),e.value.is_top?(i(),r(w,{key:0,class:"top-tag",type:"warning",size:"small",round:""},{default:s(()=>[_(" 置顶 ")]),_:1})):c("",!0),e.value.visibility==1?(i(),r(w,{key:1,class:"top-tag",type:"error",size:"small",round:""},{default:s(()=>[_(" 私密 ")]),_:1})):c("",!0),e.value.visibility==2?(i(),r(w,{key:2,class:"top-tag",type:"info",size:"small",round:""},{default:s(()=>[_(" 好友可见 ")]),_:1})):c("",!0),u("div",null,[u("span",me,m(p(K)(e.value.created_on))+" "+m(e.value.ip_loc),1)])]),"header-extra":s(()=>[u("div",de,[o($,{placement:"bottom-end",trigger:"click",size:"small",options:x.value,onSelect:O},{default:s(()=>[o(L,{quaternary:"",circle:""},{icon:s(()=>[o(p(y),null,{default:s(()=>[o(p(ie))]),_:1})]),_:1})]),_:1},8,["options"])])]),footer:s(()=>[e.value.attachments.length>0?(i(),r(b,{key:0,attachments:e.value.attachments},null,8,["attachments"])):c("",!0),e.value.charge_attachments.length>0?(i(),r(b,{key:1,attachments:e.value.charge_attachments,price:e.value.attachment_price},null,8,["attachments","price"])):c("",!0),e.value.imgs.length>0?(i(),r(j,{key:2,imgs:e.value.imgs},null,8,["imgs"])):c("",!0),e.value.videos.length>0?(i(),r(g,{key:3,videos:e.value.videos},null,8,["videos"])):c("",!0),e.value.links.length>0?(i(),r(H,{key:4,links:e.value.links},null,8,["links"])):c("",!0)]),action:s(()=>[o(se,{justify:"space-between"},{default:s(()=>[u("div",he,[o(p(y),{size:"18",class:"opt-item-icon"},{default:s(()=>[o(p(U))]),_:1}),_(" "+m(e.value.upvote_count),1)]),u("div",{class:"opt-item",onClick:n[3]||(n[3]=f(z=>a(e.value.id),["stop"]))},[o(p(y),{size:"18",class:"opt-item-icon"},{default:s(()=>[o(p(W))]),_:1}),_(" "+m(e.value.comment_count),1)]),u("div",ge,[o(p(y),{size:"18",class:"opt-item-icon"},{default:s(()=>[o(p(X))]),_:1}),_(" "+m(e.value.collection_count),1)])]),_:1})]),_:2},[e.value.texts.length>0?{name:"description",fn:s(()=>[u("div",{onClick:n[2]||(n[2]=z=>a(e.value.id))},[(i(!0),k(Q,null,R(e.value.texts,z=>(i(),k("span",{key:z.id,class:"post-text",onClick:n[1]||(n[1]=f(ne=>v(ne,e.value.id),["stop"])),innerHTML:p(B)(z.content).content},null,8,ve))),128))])]),key:"0"}:void 0]),1024)])}}});const ye={class:"nickname-wrap"},ke={class:"username-wrap"},fe={class:"item-header-extra"},xe={class:"timestamp"},we=["innerHTML"],$e={class:"opt-item"},be={class:"opt-item"},Pe=F({__name:"post-item",props:{post:{}},setup(C){const q=C,h=J(),T=G(),t=P(()=>{let e=Object.assign({texts:[],imgs:[],videos:[],links:[],attachments:[],charge_attachments:[]},q.post);return e.contents.map(a=>{(+a.type==1||+a.type==2)&&e.texts.push(a),+a.type==3&&e.imgs.push(a),+a.type==4&&e.videos.push(a),+a.type==6&&e.links.push(a),+a.type==7&&e.attachments.push(a),+a.type==8&&e.charge_attachments.push(a)}),e}),x=e=>{h.push({name:"post",query:{id:e}})},O=(e,a)=>{if(e.target.dataset.detail){const v=e.target.dataset.detail.split(":");if(v.length===2){T.commit("refresh"),v[0]==="tag"?h.push({name:"home",query:{q:v[1],t:"tag"}}):h.push({name:"user",query:{s:v[1]}});return}}x(a)};return(e,a)=>{const v=Y,l=A("router-link"),n=Z,d=N,M=V,w=D,L=I,$=y,b=ee,j=te;return i(),k("div",{class:"post-item",onClick:a[3]||(a[3]=g=>x(t.value.id))},[o(j,{"content-indented":""},E({avatar:s(()=>[o(v,{round:"",size:30,src:t.value.user.avatar},null,8,["src"])]),header:s(()=>[u("span",ye,[o(l,{onClick:a[0]||(a[0]=f(()=>{},["stop"])),class:"username-link",to:{name:"user",query:{s:t.value.user.username}}},{default:s(()=>[_(m(t.value.user.nickname),1)]),_:1},8,["to"])]),u("span",ke," @"+m(t.value.user.username),1),t.value.is_top?(i(),r(n,{key:0,class:"top-tag",type:"warning",size:"small",round:""},{default:s(()=>[_(" 置顶 ")]),_:1})):c("",!0),t.value.visibility==1?(i(),r(n,{key:1,class:"top-tag",type:"error",size:"small",round:""},{default:s(()=>[_(" 私密 ")]),_:1})):c("",!0),t.value.visibility==2?(i(),r(n,{key:2,class:"top-tag",type:"info",size:"small",round:""},{default:s(()=>[_(" 好友可见 ")]),_:1})):c("",!0)]),"header-extra":s(()=>[u("div",fe,[u("span",xe,m(t.value.ip_loc?t.value.ip_loc+" · ":t.value.ip_loc)+" "+m(p(K)(t.value.created_on)),1)])]),footer:s(()=>[t.value.attachments.length>0?(i(),r(d,{key:0,attachments:t.value.attachments},null,8,["attachments"])):c("",!0),t.value.charge_attachments.length>0?(i(),r(d,{key:1,attachments:t.value.charge_attachments,price:t.value.attachment_price},null,8,["attachments","price"])):c("",!0),t.value.imgs.length>0?(i(),r(M,{key:2,imgs:t.value.imgs},null,8,["imgs"])):c("",!0),t.value.videos.length>0?(i(),r(w,{key:3,videos:t.value.videos},null,8,["videos"])):c("",!0),t.value.links.length>0?(i(),r(L,{key:4,links:t.value.links},null,8,["links"])):c("",!0)]),action:s(()=>[o(b,{justify:"space-between"},{default:s(()=>[u("div",$e,[o($,{size:"18",class:"opt-item-icon"},{default:s(()=>[o(p(U))]),_:1}),_(" "+m(t.value.upvote_count),1)]),u("div",{class:"opt-item",onClick:a[2]||(a[2]=f(g=>x(t.value.id),["stop"]))},[o($,{size:"18",class:"opt-item-icon"},{default:s(()=>[o(p(W))]),_:1}),_(" "+m(t.value.comment_count),1)]),u("div",be,[o($,{size:"18",class:"opt-item-icon"},{default:s(()=>[o(p(X))]),_:1}),_(" "+m(t.value.collection_count),1)])]),_:1})]),_:2},[t.value.texts.length>0?{name:"description",fn:s(()=>[(i(!0),k(Q,null,R(t.value.texts,g=>(i(),k("span",{key:g.id,class:"post-text",onClick:a[1]||(a[1]=f(H=>O(H,t.value.id),["stop"])),innerHTML:p(B)(g.content).content},null,8,we))),128))]),key:"0"}:void 0]),1024)])}}});export{Pe as _,He as a};
